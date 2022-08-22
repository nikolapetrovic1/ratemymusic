import { Component, OnInit } from '@angular/core';
import { MusicianService } from '../../service/musician.service';

@Component({
  selector: 'app-all-musicians',
  templateUrl: './all-musicians.component.html',
  styleUrls: ['./all-musicians.component.scss']
})
export class AllMusiciansComponent implements OnInit {


  query: string = '';
  musicians: any = [];
  currentRate = 1;
  constructor(private musicianService:MusicianService) { }

  ngOnInit(): void {
    this.search();
  }

  search(){
    this.musicianService.search(this.query).subscribe((res)=>{
      this.musicians = res.body.musicians;
    })
  }

}
