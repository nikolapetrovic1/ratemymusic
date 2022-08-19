import { Component, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';
import { ActivatedRoute } from '@angular/router';
import { FormControl } from '@angular/forms';
import { RatingService } from '../../service/rating.service';
@Component({
  selector: 'app-song-profile',
  templateUrl: './song-profile.component.html',
  styleUrls: ['./song-profile.component.scss']
})
export class SongProfileComponent implements OnInit {
  song!: any;
  ratingValue = new FormControl('');
  songId!: any;
  ratings! : any;



  constructor(private router: ActivatedRoute,private songService:SongService,private ratingService:RatingService) { }

  ngOnInit(): void {
    this.songId = JSON.parse(this.router.snapshot.paramMap.get('id')!);
    this.loadRatings();
  }
  getAllSongRatings(songId:string){
    this.ratingService.getAllSongRatings(songId).subscribe((res)=>{
      this.ratings = res.rating_data;
    })
  }
  getUserRating(songId:string){
    this.
    ratingService.getUserRating(songId).subscribe((res)=>{
      console.log(res);
      this.ratingValue.setValue(res.rating_value);
    })
  }
  getSongId(id:string){
    this.songService.get(id).subscribe((res)=>{
        this.song = res.data;
    })
  }
  rateSong(){
    var rating : any = {
      rating_value : this.ratingValue.value,
      song_id : this.song.id,
    }
    this.ratingService.rate(rating).subscribe((res)=>{
      this.loadRatings();
    })
  }
  loadRatings(){
    this.getSongId(this.songId)
    this.getUserRating(this.songId);
    this.getAllSongRatings(this.songId);
  }

}
