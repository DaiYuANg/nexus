use crate::app_state::AppState;
use crate::config::Config;
use crate::entry::EntryService;
use crate::entry::http::HttpService;
use crate::entry::tcp::TcpService;
use figment::Figment;
use figment::providers::{Env, Format, Serialized, Toml};
use tracing::{error, info};

mod app_state;
mod config;
mod entry;
mod indexer;
mod repository;
mod storage;

fn init() -> Config {
  tracing_subscriber::fmt()
    .with_max_level(tracing::Level::DEBUG)
    .init();
  dotenvy::dotenv().ok();
  let config: Config = Figment::new()
    .merge(Serialized::defaults(Config {
      debug: false,
      port: 8080,
    }))
    .merge(Toml::file("config.toml").nested())
    .merge(Env::prefixed("STORIX_"))
    .extract()
    .expect("Failed to load config");

  config
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
  let _config = init();

  let state = AppState {};

  let services: Vec<Box<dyn EntryService>> = vec![
    Box::new(HttpService {
      addr: "0.0.0.0:8080".parse()?,
      app_state: state.clone(),
    }),
    Box::new(TcpService {
      addr: "0.0.0.0:9090".parse()?,
      app_state: state.clone(),
    }),
  ];

  let handles = services.into_iter().map(|svc| {
    tokio::spawn(async move {
      info!("Starting service: {}", svc.name());
      if let Err(e) = svc.start().await {
        error!("Service {} failed: {:?}", svc.name(), e);
      }
    })
  });
  futures::future::join_all(handles).await;
  Ok(())
}
