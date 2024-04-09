use serde::{Deserialize, Serialize};
use warp::Filter;
use rdkafka::config::ClientConfig;
use rdkafka::producer::{FutureProducer, FutureRecord};
use rdkafka::util::Timeout;

#[derive(Deserialize, Serialize)]
struct Data {
    album: String,
    year: i32,
    artist: String,
    ranked: i32,
}

async fn insert_data(data: Data) -> Result<impl warp::Reply, warp::Rejection> {
    let message = serde_json::to_string(&data).expect("Can serialize data");

    // Configuración y creación del productor de Kafka
    let producer: FutureProducer = ClientConfig::new()
        .set("bootstrap.servers", "kafka:9092")
        .set("message.timeout.ms", "5000")
        .create()
        .expect("Producer creation error");

    // Creación y envío del mensaje a Kafka
    let record = FutureRecord::to("vote-topic")
        .payload(&message)
        .key(&data.album);

    match producer.send(record, Timeout::Never).await {
        Ok(_) => Ok(warp::reply::json(&"Mensaje recibido y enviado a Kafka")),
        Err(e) => Err(warp::reject::custom(e)),
    }
}

#[tokio::main]
async fn main() {
    let insert_route = warp::post()
        .and(warp::path("insert"))
        .and(warp::body::json())
        .and_then(insert_data);

    warp::serve(insert_route)
        .run(([0, 0, 0, 0], 8081))
        .await;
}
