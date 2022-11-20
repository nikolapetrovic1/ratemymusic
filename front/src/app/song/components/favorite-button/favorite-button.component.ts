import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { CollectionService } from 'src/app/collection/service/collection.service';

@Component({
  selector: 'app-favorite-button',
  templateUrl: './favorite-button.component.html',
  styleUrls: ['./favorite-button.component.scss']
})
export class FavoriteButtonComponent implements OnInit {
  @Input() kind: any;
  @Input() kindId: any;
  @Input() name: any;
  @Output() changed = new EventEmitter();


  favorite: any;

  constructor(private collectionService:CollectionService) { }

  ngOnInit(): void {
    this.loadFavorite();
  }

  loadFavorite(){
    this.collectionService.getFavoriteUserKindId(this.kind,this.kindId).subscribe((res)=>{
      this.favorite = res.favorite;
      console.log(this.favorite);
    })
  }
  changeFavorite(){
    if(this.favorite){
      this.collectionService.unFavorite(this.favorite.id).subscribe((res)=>{
        this.loadFavorite();
        console.log(res);
      })
    }
    else{
      var favoriteJson = {
        kind: this.kind,
        kindId: this.kindId,
        name: this.name
      }
      this.collectionService.favorite(favoriteJson).subscribe((res)=>{
        this.loadFavorite();
        console.log(res);
      })
    }

  }

}
