import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateMusicianComponent } from './create-musician.component';

describe('CreateMusicianComponent', () => {
  let component: CreateMusicianComponent;
  let fixture: ComponentFixture<CreateMusicianComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateMusicianComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateMusicianComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
