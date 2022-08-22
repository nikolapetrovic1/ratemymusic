use crate::comment::CommentRequest;
use crate::models::{Comment, NewComment};
use crate::db::establish_connection;
use diesel::prelude::*;