use crate::{schema::comments, comment::{CommentRequest}};


#[derive(Insertable)]
#[table_name = "comments"]
pub struct NewComment<'a> {
    pub comment:&'a str,
    pub user_id: i32,
    pub review_id: i32,
}

#[derive(Debug, Queryable, AsChangeset)]
pub struct Comment {
    pub id: i32,
    pub comment:String,
    pub user_id: i32,
    pub review_id: i32,
}

pub fn map_comment_to_comment_request(comment_value:&Comment) -> CommentRequest{
    let comment_request = CommentRequest {
        id: comment_value.id,
        comment: comment_value.comment.clone(),
        user_id: comment_value.user_id,
        review_id: comment_value.review_id,
    };
    return comment_request;
}
