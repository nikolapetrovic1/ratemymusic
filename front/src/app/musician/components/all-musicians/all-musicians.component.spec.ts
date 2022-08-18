import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllMusiciansComponent } from './all-musicians.component';

describe('AllMusiciansComponent', () => {
  let component: AllMusiciansComponent;
  let fixture: ComponentFixture<AllMusiciansComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllMusiciansComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllMusiciansComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
