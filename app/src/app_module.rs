use nject::{injectable, provider};

#[injectable]
pub struct DepOne;

#[injectable]
pub struct DepTwo {
  dep: DepOne,
}

#[injectable]
pub struct Facade {
  dep: DepTwo,
}

#[provider]
pub struct AppProvider {}
