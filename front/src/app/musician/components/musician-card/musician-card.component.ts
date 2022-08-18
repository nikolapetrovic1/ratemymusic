import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-musician-card',
  templateUrl: './musician-card.component.html',
  styleUrls: ['./musician-card.component.scss']
})
export class MusicianCardComponent implements OnInit {


  @Input() musician: any;
  constructor() { }

  ngOnInit(): void {
  }

}
