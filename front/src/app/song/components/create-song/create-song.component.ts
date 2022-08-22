import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Observable } from 'rxjs';
import { AlbumService } from 'src/app/album/service/album.service';
import { MusicianService } from 'src/app/musician/service/musician.service';
import { SongService } from '../../service/song.service';
import {startWith, map} from 'rxjs/operators';
@Component({
  selector: 'app-create-song',
  templateUrl: './create-song.component.html',
  styleUrls: ['./create-song.component.scss']
})
export class CreateSongComponent implements OnInit {
  form!: FormGroup;
  albums!: any;
  filteredAlbums!: Observable<any[]>;
  musicians!: any;
  filteredMusicians!: Observable<any[]>;
  selectedAlbum= new FormControl<any>('');
  selectedMusician = new FormControl<any>('');
  constructor(private fb:FormBuilder,private songService:SongService,private albumService:AlbumService,private musicianService:MusicianService) {

    this.form = this.fb.group({
      name : ['',Validators.required],
      musician_id: ['',Validators.required],
      album_id:  [''],
      duration: [0,Validators.required],
    })
    this.loadMusicians();
   }

  ngOnInit(): void {
  }

  submit(){
    this.songService.submit(this.form.value).subscribe((res)=>{
      console.log(res)
    })
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
  displayFn(musician: any): string {
    return musician.musician_name;
  }

  private _filterMusicians(value: string): any[] {
    const filterValue = value;

    return this.musicians.filter(((musician: { musician_name: string, musician_id:string }) => musician.musician_name.toLowerCase().includes(filterValue)));
  }
  private _filterAlbums(value: string): any[] {
    const filterValue = value;

    return this.albums.filter(((album: { name: string}) => album.name.toLowerCase().includes(filterValue)));
  }
  loadAlbums(id:any){
    this.albumService.loadByMusician(id).subscribe((res)=>{
      this.albums = res.albums
      this.filteredAlbums = this.form.get('album_id')!.valueChanges.pipe(
        startWith(''),
          map(value => {
          const name = typeof value === 'string' ? value : value?.name;
          return name ? this._filterAlbums(name as string) : this.albums.slice();
        }),
      );
    })
  }

}
