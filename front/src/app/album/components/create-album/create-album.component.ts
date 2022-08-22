import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { AlbumService } from '../../service/album.service';
import {startWith, map} from 'rxjs/operators';
import { MusicianService } from 'src/app/musician/service/musician.service';
@Component({
  selector: 'app-create-album',
  templateUrl: './create-album.component.html',
  styleUrls: ['./create-album.component.scss']
})
export class CreateAlbumComponent implements OnInit {

  form!: FormGroup;
  musicians!: any;
  filteredMusicians!: Observable<any[]>;
  myControl = new FormControl<any>('');
  constructor(private fb:FormBuilder,private albumService:AlbumService,private musicianService:MusicianService) {

    this.form = this.fb.group({
      name : ['',Validators.required],
      musician_id: [''],
    })
    this.loadMusicians();
   }
  loadMusicians(){
    this.musicianService.search("").subscribe((res)=>{
      this.musicians = res.body.musicians;
      this.filteredMusicians = this.form.get('musician_id')!.valueChanges.pipe(
        startWith(''),
          map(value => {
          const name = typeof value === 'string' ? value : value?.name;
          return name ? this._filterMusicians(name as string) : this.musicians.slice();
        }),
      );
    })
  }
  ngOnInit(): void {
  }
  submit(){ 
    this.form.controls['musician_id'].setValue(this.myControl.value.id)
    this.albumService.submit(this.form.value).subscribe((res)=>{
      console.log(res);
    })
  }
  displayFn(musician: any): string {
    return musician.musician_name;
  }
  private _filterMusicians(value: string): any[] {
    const filterValue = value;

    return this.musicians.filter(((musician: { musician_name: string, musician_id:string }) => musician.musician_name.toLowerCase().includes(filterValue)));
  }
}
