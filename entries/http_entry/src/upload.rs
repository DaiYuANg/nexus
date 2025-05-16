use axum::extract::Multipart;

pub async fn accept_form(mut multipart: Multipart) {
  while let Some(field) = multipart.next_field().await.unwrap() {
    let name = field.name().unwrap().to_string();
    let file_name = field.file_name().unwrap().to_string();
    let content_type = field.content_type().unwrap().to_string();
    let data = field.bytes().await.unwrap();

    println!(
      "Length of `{name}` (`{file_name}`: `{content_type}`) is {} bytes",
      data.len()
    );
  }
}