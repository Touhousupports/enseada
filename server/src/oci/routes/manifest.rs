use actix_web::web::{Data, Json, Path};
use actix_web::HttpResponse;
use serde::Deserialize;

use crate::couchdb::repository::Repository;
use crate::oci::digest::Digest;
use crate::oci::entity::{Manifest, Repo};
use crate::oci::error::{Error, ErrorCode};
use crate::oci::manifest::ImageManifest;
use crate::oci::mime::MediaType;
use crate::oci::routes::RepoPath;
use crate::oci::service::{ManifestService, RepoService};
use crate::oci::{header, Result};
use std::convert::TryFrom;

#[derive(Debug, Deserialize)]
pub struct ManifestRefParam {
    reference: String,
}

pub async fn get(
    manifests: Data<ManifestService>,
    repos: Data<RepoService>,
    repo: Path<RepoPath>,
    reference: Path<ManifestRefParam>,
) -> Result<HttpResponse> {
    let group = &repo.group;
    let name = &repo.name;
    let reference = &reference.reference;

    log::debug!("looking for repo {}/{}", group, name);
    repos
        .find(&Repo::build_id(group, name))
        .await?
        .ok_or_else(|| Error::from(ErrorCode::NameUnknown))?;

    let manifest = manifests
        .find_by_ref(reference)
        .await?
        .ok_or_else(|| Error::from(ErrorCode::ManifestUnknown))?;

    let manifest = manifest.into_inner();
    Ok(HttpResponse::Ok()
        .header(http::header::CONTENT_TYPE, MediaType::ImageManifest)
        .header(header::CONTENT_DIGEST, manifest.digest().to_string())
        .json(manifest))
}

pub async fn put(
    manifests: Data<ManifestService>,
    repos: Data<RepoService>,
    repo: Path<RepoPath>,
    reference: Path<ManifestRefParam>,
    body: Json<ImageManifest>,
) -> Result<HttpResponse> {
    let group = &repo.group;
    let name = &repo.name;
    let reference = &reference.reference;

    log::debug!("looking for repo {}/{}", group, name);
    let repo = repos
        .find(&Repo::build_id(group, name))
        .await?
        .ok_or_else(|| Error::from(ErrorCode::NameUnknown))?;

    let manifest = Manifest::new(reference, body.into_inner());
    let manifest = manifests.save(manifest).await?;

    log::debug!("Checking if ref '{}' is a tag", reference);
    if Digest::try_from(reference).is_err() {
        log::debug!("Ref '{}' is indeed a tag", reference);
        // Reference is a tag
        let mut repo = repo;
        repo.push_tag(reference.clone());
        repos.save(repo).await?;
    }

    let manifest = manifest.into_inner();

    Ok(HttpResponse::Created()
        .header(
            http::header::LOCATION,
            format!("/{}/{}/manifests/{}", group, name, reference),
        )
        .header(header::CONTENT_DIGEST, manifest.digest().to_string())
        .finish())
}

pub async fn delete(
    manifests: Data<ManifestService>,
    repos: Data<RepoService>,
    repo: Path<RepoPath>,
    reference: Path<ManifestRefParam>,
) -> Result<HttpResponse> {
    let group = &repo.group;
    let name = &repo.name;
    let reference = &reference.reference;

    log::debug!("looking for repo {}/{}", group, name);
    repos
        .find(&Repo::build_id(group, name))
        .await?
        .ok_or_else(|| Error::from(ErrorCode::NameUnknown))?;

    let manifest = manifests
        .find_by_ref(reference)
        .await?
        .ok_or_else(|| Error::from(ErrorCode::ManifestUnknown))?;

    manifests.delete(&manifest).await?;

    Ok(HttpResponse::Accepted().finish())
}
