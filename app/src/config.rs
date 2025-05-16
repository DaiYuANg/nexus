use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug, Serialize, Clone)]
pub struct Config {
  pub(crate) debug: bool,
  pub(crate) port: u16,
  pub(crate) http_config: HttpConfig,
}

#[derive(Deserialize, Debug, Serialize, Clone)]
pub struct HttpConfig {
  pub(crate) port: u16,
  pub(crate) host: String,
}

impl HttpConfig {
  pub fn listen_address(&self) -> String {
    format!("{}:{}", self.host, self.port)
  }
}
