use wasm_bindgen::prelude::*;
use web_sys::{console, Request, RequestInit, RequestMode, Response};
use serde::{Serialize, Deserialize};
use wasm_bindgen_futures::JsFuture;
use js_sys::Promise;

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
        album: "Test Album".into(),
        year: "2022".into(),
        artist: "Test Artist".into(),
        ranked: "1".into(),
    };

    let mut opts = RequestInit::new();
    opts.method("POST");
    opts.mode(RequestMode::Cors);
    let request_body = serde_json::to_string(&data).unwrap();
    opts.body(Some(&JsValue::from_str(&request_body)));

    let request = Request::new_with_str_and_init("/insert", &opts)?;
    request.headers().set("Content-Type", "application/json")?;

    let window = web_sys::window().unwrap();
    let resp_value = JsFuture::from(window.fetch_with_request(&request)).await?;
    let resp: Response = resp_value.dyn_into().unwrap();
    console::log_1(&resp.text().await?.into());

    Ok(())
}
