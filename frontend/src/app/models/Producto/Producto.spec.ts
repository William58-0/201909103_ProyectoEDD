import { Producto } from './Producto';

describe('Producto', () => {
  it('should create an instance', () => {
    expect(new Producto("",0,"","",0,"","","","",0)).toBeTruthy();
  });
});