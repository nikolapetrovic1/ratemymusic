import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MusicianService } from '../../service/musician.service';

@Component({
  selector: 'app-create-musician',
  templateUrl: './create-musician.component.html',
  styleUrls: ['./create-musician.component.scss']
})
export class CreateMusicianComponent implements OnInit {
  form!:FormGroup;

  constructor(private fb:FormBuilder,private musicianService:MusicianService) {
    this.form = this.fb.group({
      musician_name : ['',Validators.required],
      name: ['',Validators.required],
      surname : ['',Validators.required],
    })
   }

  ngOnInit(): void {
  }
  submit(){
    this.musicianService.submit(this.form.value).subscribe((res)=>{
      console.log(res)
    })
  }
}
