///|
pub struct Path {
  checkpoints : FixedArray[Vec2]
  next : Vec2
} derive(Default)

///|
pub fn Path::new(checkpoints : FixedArray[Vec2]) -> Path {
  Path::{ checkpoints, next: checkpoints.get(0).unwrap_or_default() }
}
