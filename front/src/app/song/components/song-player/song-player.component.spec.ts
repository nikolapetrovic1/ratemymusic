import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SongPlayerComponent } from './song-player.component';

describe('SongPlayerComponent', () => {
  let component: SongPlayerComponent;
  let fixture: ComponentFixture<SongPlayerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SongPlayerComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SongPlayerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
