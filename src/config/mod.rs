use serde::{Deserialize, Serialize};

#[derive(Deserialize, Debug, Serialize)]
pub struct Config {
  pub(crate) debug: bool,
  pub(crate) port: u16,
}
