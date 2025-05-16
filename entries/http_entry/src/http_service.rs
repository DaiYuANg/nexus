use crate::upload;
use axum::extract::DefaultBodyLimit;
use axum::routing::{get, post};
use entry::entry_service::EntryService;
use metrics_exporter_prometheus::PrometheusBuilder;
use std::net::SocketAddr;
use tower_http::limit::RequestBodyLimitLayer;
use tracing::log::info;
use utoipa::OpenApi;
use utoipa_axum::router::OpenApiRouter;

// state you need
#[derive(Clone)]
pub(crate) struct AppState {}

pub struct HttpService {
  pub addr: SocketAddr,
}
const TAG: &str = "ROOT";
#[derive(OpenApi)]
#[openapi(
  tags(
        (name = TAG, description = "Storix API endpoints")
  )
)]
struct ApiDoc;

#[async_trait::async_trait]
impl EntryService for HttpService {
  async fn start(&self) -> anyhow::Result<()> {
    let builder = PrometheusBuilder::new();
    let recorder = builder
      .install_recorder()
      .expect("failed to install recorder");
    let state = AppState {};
    let (router, api) = OpenApiRouter::with_openapi(ApiDoc::openapi())
      .with_state(state)
      .layer(DefaultBodyLimit::disable())
      .layer(tower_http::trace::TraceLayer::new_for_http())
      .layer(RequestBodyLimitLayer::new(250 * 1024 * 1024))
      .route("/", get(|| async { "Hello, FileDB!" }))
      .route(
        "/metrics",
        get(move || {
          let body = recorder.render();
          async move { axum::response::Html(body) }
        }),
      )
      .route("/upload", post(upload::accept_form))
      .split_for_parts();
    let listener = tokio::net::TcpListener::bind(self.addr).await?;
    info!(
      "{}",
      "🚀 Starting storix HTTP server at http://127.0.0.1:8080".to_string()
    );
    axum::serve(listener, router).await?;
    Ok(())
  }

  fn name(&self) -> &'static str {
    "http"
  }
}
