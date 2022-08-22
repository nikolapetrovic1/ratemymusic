import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-song-card',
  templateUrl: './song-card.component.html',
  styleUrls: ['./song-card.component.scss']
})
export class SongCardComponent implements OnInit {

  @Input() song: any;
  constructor(private router:Router) { }

  ngOnInit(): void {
  }
  songPage(id:any){
    this.router.navigate([`/song/${id}`]);
  }
}
