import { ComponentFixture, TestBed } from '@angular/core/testing';

import { MusicianProfileComponent } from './musician-profile.component';

describe('MusicianProfileComponent', () => {
  let component: MusicianProfileComponent;
  let fixture: ComponentFixture<MusicianProfileComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ MusicianProfileComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(MusicianProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
