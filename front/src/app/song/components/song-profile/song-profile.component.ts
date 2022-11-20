import { Component, OnInit } from '@angular/core';
import { SongService } from '../../service/song.service';
import { ActivatedRoute } from '@angular/router';
import { FormControl } from '@angular/forms';
import { RatingService } from '../../service/rating.service';
import { MusicianService } from 'src/app/musician/service/musician.service';
import { AuthService } from 'src/app/auth/service/auth.service';
import { CollectionService } from 'src/app/collection/service/collection.service';
@Component({
  selector: 'app-song-profile',
  templateUrl: './song-profile.component.html',
  styleUrls: ['./song-profile.component.scss']
})
export class SongProfileComponent implements OnInit {
  song!: any;
  ratingValue = new FormControl('');
  songId!: any;
  ratings!: any;
  averageRating!: number;
  musician!: any;
  favorite!: any;
  constructor(
    private router: ActivatedRoute,
    private songService: SongService,
    private ratingService: RatingService,
    private musicianService: MusicianService,
    private authService: AuthService,
    ) { }

  ngOnInit(): void {
    this.songId = JSON.parse(this.router.snapshot.paramMap.get('id')!);
    this.loadRatings();
  }

  getAllSongRatings(songId: string) {
    this.ratingService.getAllSongRatings(songId, 'song').subscribe((res) => {
      this.ratings = res.rating_data;
      if (this.ratings) {
        this.calculateAverage();
      }

    })
  }
  getUserRating(songId: string) {
    this.ratingService.getUserRating(songId, 'song').subscribe((res) => {
      this.ratingValue.setValue(res.rating_value);
    })
  }
  getSongId(id: string) {
    this.songService.get(id).subscribe((res) => {
      this.song = res.data;
      this.getMusicianInfo(this.song.musicianID)

    })
  }
  getMusicianInfo(id: number) {
    this.musicianService.getById(id).subscribe((res) => {
      this.musician = res.data;
    })
  }
  rateSong() {
    var rating: any = {
      rating_value: this.ratingValue.value,
      song_id: this.song.id,
    }
    this.ratingService.rate(rating, "song").subscribe((res) => {
      this.loadRatings();
    })
  }
  loadRatings() {
    this.getSongId(this.songId)
    this.getUserRating(this.songId);
    this.getAllSongRatings(this.songId);
  }
  calculateAverage() {
    this.averageRating = 0;
    this.ratings.forEach((rating: { rating_value: number; }) => { this.averageRating += rating.rating_value });
    this.averageRating = this.averageRating / this.ratings.length
  }
  isLoggedIn() {
    return this.authService.isLoggedIn();
  }
}
