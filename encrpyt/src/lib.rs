use anyhow::Result;
use serde::{Deserialize, Serialize};
use spin_sdk::{
    http::{Request, Response},
    http_component,
};
use magic_crypt::{new_magic_crypt, MagicCryptTrait};

#[derive(Serialize, Deserialize, Debug)]
struct RequestPayload {
    message: String,
}

#[derive(Serialize, Deserialize, Debug)]
struct ResponsePayload {
    message: String,
}

/// A simple Spin HTTP component.
#[http_component]
fn encrpyt(req: Request) -> Result<Response> {
    let request_payload: RequestPayload =
    match serde_json::from_slice(req.body().clone().unwrap().to_vec().as_slice()) {
        Ok(request_payload) => request_payload,
        Err(_) => {
            return Ok(http::Response::builder()
                .status(500)
                .body(Some("Failed to parse payload".into()))?);
        }
    };

    let mc = new_magic_crypt!("magickey", 256);
    let base64 = mc.encrypt_str_to_base64(&request_payload.message);

    println!("{:?}", request_payload);
    println!("{:?}", base64);

    let response: ResponsePayload = ResponsePayload { 
        message: base64
    };

    Ok(http::Response::builder()
        .status(200)
        .body(Some(serde_json::to_string(&response).unwrap().into()))?)
}
