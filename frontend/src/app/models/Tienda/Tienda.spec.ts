import { Tienda } from './Tienda';

describe('Tienda', () => {
  it('should create an instance', () => {
    expect(new Tienda("","","",0,"","")).toBeTruthy();
  });
});
