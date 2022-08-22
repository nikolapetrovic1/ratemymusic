import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-musician-card',
  templateUrl: './musician-card.component.html',
  styleUrls: ['./musician-card.component.scss']
})
export class MusicianCardComponent implements OnInit {


  @Input() musician: any;
  constructor(private router:Router) { }

  ngOnInit(): void {
  }

  musicianPage(id:any){
    this.router.navigate([`/musician/${id}`]);
  }
}
