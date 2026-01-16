///|
using @firefly {
  type Point,
  type Angle,
  type Size,
  type Style,
  type LineStyle,
  type Pad,
}

///|
let world : Ref[World] = Ref::new(World::default())

///|
fn main {

}

///|
pub fn boot() -> Unit {
  let assets = Assets::load()
  world.val = World::{ ..world.val }
}

///|
pub fn update() -> Unit {
  world.val.update()
}

///|
pub fn render() -> Unit {
  world.val.render()
}
