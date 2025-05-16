// src/entry/http.rs
use super::EntryService;
use crate::app_state::AppState;
use axum::{Router, routing::get};
use metrics_exporter_prometheus::PrometheusBuilder;
use std::net::SocketAddr;
use tracing::log::info;

pub struct HttpService {
  pub addr: SocketAddr,
  pub app_state: AppState,
}

#[async_trait::async_trait]
impl EntryService for HttpService {
  async fn start(&self) -> anyhow::Result<()> {
    let builder = PrometheusBuilder::new();
    let recorder = builder
      .install_recorder()
      .expect("failed to install recorder");
    let state = self.app_state.clone();
    let app = Router::new()
      .with_state(state)
      .route("/", get(|| async { "Hello, FileDB!" }))
      .route(
        "/metrics",
        get(move || {
          let body = recorder.render();
          async move { axum::response::Html(body) }
        }),
      );
    let listener = tokio::net::TcpListener::bind(self.addr).await?;
    info!("🚀 Starting storix HTTP server at http://127.0.0.1:8080");
    axum::serve(listener, app).await?;
    Ok(())
  }

  fn name(&self) -> &'static str {
    "http"
  }
}
