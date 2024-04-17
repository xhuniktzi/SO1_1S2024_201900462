use wasm_bindgen::prelude::*;
use wasm_bindgen_futures::JsFuture;
use web_sys::{console, Headers, Request, RequestInit, Response};

use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct Data {
    album: String,
    year: String,
    artist: String,
    ranked: String,
}

#[wasm_bindgen(start)]
pub async fn run() -> Result<(), JsValue> {
    let data = Data {
        album: "Test Album".to_string(),
        year: "2022".to_string(),
        artist: "Test Artist".to_string(),
        ranked: "1".to_string(),
    };

    let mut opts = RequestInit::new();
    opts.method("POST");
    let request_body = serde_json::to_string(&data).unwrap();
    opts.body(Some(&JsValue::from_str(&request_body)));

    let request = Request::new_with_str_and_init("/insert", &opts)?;
    let headers = Headers::new()?;
    headers.set("Content-Type", "application/json")?;
    request.headers().set("Content-Type", "application/json")?;

    let window = web_sys::window().unwrap();
    let resp_value = JsFuture::from(window.fetch_with_request(&request)).await?;
    let resp: Response = resp_value.dyn_into().unwrap();

    // Convert the Promise returned by resp.text() into a Future and await it
    let text_promise = resp.text()?;
    let text_jsvalue = JsFuture::from(text_promise).await?;
    let text = text_jsvalue.as_string().unwrap(); // Convert JsValue to String

    console::log_1(&text.into());

    Ok(())
}
