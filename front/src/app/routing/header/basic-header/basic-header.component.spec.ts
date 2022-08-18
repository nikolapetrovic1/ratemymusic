import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BasicHeaderComponent } from './basic-header.component';

describe('BasicHeaderComponent', () => {
  let component: BasicHeaderComponent;
  let fixture: ComponentFixture<BasicHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BasicHeaderComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(BasicHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
