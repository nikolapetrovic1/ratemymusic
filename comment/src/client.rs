use comment::comment_client::CommentClient;
use comment::CommentRequest;

pub mod comment {
    tonic::include_proto!("comment");
}
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = CommentClient::connect("http://[::1]:50055").await?;

    let request = tonic::Request::new(
        CommentRequest { comment: String::from("Test"), user_id: 1, review_id: 1 }
    );

    let response = client.create_comment(request).await?;

    println!("RESPONSE={:?}", response);

    Ok(())
}