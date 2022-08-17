use crate::{schema::comments, comment::CommentRequest};


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


// impl From<CommentRequest> for NewComment<'_> {
//     fn from(request: CommentRequest) -> Self {
//       Self {
//        comment : &request.comment,
//        user_id : request.user_id,
//        review_id : request.review_id,
//       }
//     }
//   }