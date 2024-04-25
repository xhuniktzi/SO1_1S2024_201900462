use rocket::response::status::BadRequest;
use rocket::serde::json::{json, Value as JsonValue};
use rocket::serde::json::Json;
use rocket::config::SecretKey;
use rocket_cors::{AllowedOrigins, CorsOptions};
use tiny_kafka::producer::{KafkaProducer, Message};

#[derive(rocket::serde::Deserialize)]
struct Data {
    artist: String,
    album: String,
    year: i32,
    ranked: i32,
}

#[rocket::post("/insert", data = "<data>")]
async fn receive_data(data: Json<Data>) -> Result<Json<JsonValue>, BadRequest<String>> {
    let received_data = data.into_inner();
    
    let response = JsonValue::from(json!({
        "message": format!("Received data: Name: {}, Album: {}, Year: {}, Rank: {}", received_data.artist, received_data.album, received_data.year, received_data.ranked)
    }));

    let msg_obj = JsonValue::from(json!({
        "artist": received_data.artist,
        "album": received_data.album,
        "year": received_data.year,
        "ranked": received_data.ranked,
    }));
    let msg_obj_string = msg_obj.to_string();
    let msg = Message::new("msg", &msg_obj_string);
    let producer = KafkaProducer::new("kafka:9092", None);
    producer.send_message("vote-topic2", msg).await;

    Ok(Json(response))
}

#[rocket::main]
async fn main() {
    let secret_key = SecretKey::generate(); // Genera una nueva clave secreta

    // Configuración de opciones CORS
    let cors = CorsOptions::default()
        .allowed_origins(AllowedOrigins::all())
        .to_cors()
        .expect("failed to create CORS fairing");

    let config = rocket::Config {
        address: "0.0.0.0".parse().unwrap(),
        port: 8081,
        secret_key: secret_key.unwrap(), // Desempaqueta la clave secreta generada
        ..rocket::Config::default()
    };

    // Montar la aplicación Rocket con el middleware CORS
    rocket::custom(config)
        .attach(cors)
        .mount("/", rocket::routes![receive_data])
        .launch()
        .await
        .unwrap();
}