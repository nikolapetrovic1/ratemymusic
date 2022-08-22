import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReviewViewComponent } from './review-view.component';

describe('ReviewViewComponent', () => {
  let component: ReviewViewComponent;
  let fixture: ComponentFixture<ReviewViewComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReviewViewComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReviewViewComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
