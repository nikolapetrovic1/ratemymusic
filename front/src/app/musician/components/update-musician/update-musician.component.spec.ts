import { ComponentFixture, TestBed } from '@angular/core/testing';

import { UpdateMusicianComponent } from './update-musician.component';

describe('UpdateMusicianComponent', () => {
  let component: UpdateMusicianComponent;
  let fixture: ComponentFixture<UpdateMusicianComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ UpdateMusicianComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(UpdateMusicianComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
