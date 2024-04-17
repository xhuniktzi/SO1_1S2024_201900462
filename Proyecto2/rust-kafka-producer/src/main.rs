use serde::{Serialize, Deserialize};
use warp::{Filter, reject, Rejection, Reply};
use rdkafka::{config::ClientConfig, producer::{FutureProducer, FutureRecord}};
use std::time::Duration;

#[derive(Debug, Deserialize, Serialize)]
struct Data {
    album: String,
    year: String,
    artist: String,
    ranked: String,
}

#[derive(Debug)]
struct KafkaErrorWrapper {
    inner: rdkafka::error::KafkaError,
}

impl reject::Reject for KafkaErrorWrapper {}

async fn insert_data(data: Data, producer: FutureProducer) -> Result<impl Reply, Rejection> {
    let message = serde_json::to_string(&data).unwrap();
    let record = FutureRecord::<(), _>::to("vote-topic")
        .payload(&message);
    match producer.send(record, Duration::from_secs(5)).await {
        Ok(_) => Ok(warp::reply::json(&data)),
        Err((e, _)) => Err(reject::custom(KafkaErrorWrapper { inner: e })),
    }
}

#[tokio::main]
async fn main() {
    let producer: FutureProducer = ClientConfig::new()
        .set("bootstrap.servers", "kafka:9092")
        .create().unwrap();

    let insert_route = warp::post()
        .and(warp::path("insert"))
        .and(warp::body::json())
        .and(warp::any().map(move || producer.clone()))
        .and_then(insert_data);

    warp::serve(insert_route)
        .run(([0, 0, 0, 0], 50051))
        .await;
}
