///|
pub struct Path {
  checkpoints : ArrayView[Vec2]
  mut next : (Int, Vec2)
}

///|
pub fn Path::new(checkpoints : ArrayView[Vec2]) -> Path {
  let path = Path::{ checkpoints, next: (-1, Vec2::new(0, 0)) }
  // TODO: uncommenting this causes compilation issue
  //path.next()
  path
}

///|
pub fn Path::next(self : Path) -> Unit {
  let idx = (self.next.0 + 1) % self.checkpoints.length()
  self.next = (idx, self.checkpoints[idx])
}

///|
pub fn Path::clone(self : Path) -> Path {
  Path::{ checkpoints: self.checkpoints, next: self.next }
}
