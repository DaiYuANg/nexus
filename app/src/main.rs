use crate::app_module::{AppProvider, Facade};
use crate::cli::Args;
use crate::config::{Config, HttpConfig};
use clap::Parser;
use entry::entry_service::EntryService;
use figment::providers::{Env, Format, Serialized, Toml};
use figment::Figment;
use http_entry::http_service::HttpService;
use nject::{injectable, provider};
use tcp_entry::tcp_service::TcpService;
use tracing::log::{error, info};
use tracing_subscriber::layer::SubscriberExt;
use tracing_subscriber::util::SubscriberInitExt;

mod app_module;
mod cli;
mod config;

fn init() -> Config {
  tracing_subscriber::registry()
    .with(
      tracing_subscriber::EnvFilter::try_from_default_env()
        .unwrap_or_else(|_| format!("{}=debug,tower_http=debug", env!("CARGO_CRATE_NAME")).into()),
    )
    .with(tracing_subscriber::fmt::layer())
    .init();
  dotenvy::dotenv().ok();
  let config: Config = Figment::new()
    .merge(Serialized::defaults(Config {
      debug: false,
      port: 8080,
      http_config: HttpConfig {
        port: 7280,
        host: "0.0.0.0".parse().unwrap(),
      },
    }))
    .merge(Toml::file("config.toml").nested())
    .merge(Env::prefixed("STORIX_"))
    .extract()
    .expect("Failed to load config");

  config
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
  let args = Args::parse();
  let _config = init();
  let _facade: Facade = AppProvider::provide(&AppProvider {});
  // let state = AppState { config: _config };
  let services: Vec<Box<dyn EntryService>> = vec![
    Box::new(HttpService {
      addr: "0.0.0.0:7820".parse()?,
    }),
    Box::new(TcpService {
      addr: "0.0.0.0:9090".parse()?,
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
