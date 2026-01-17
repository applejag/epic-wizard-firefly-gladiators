///|
using @firefly {type Size}

///|
using @util {type Vec2}

///|
let racebattle_world : Ref[@racebattle.World] = Ref::new(
  @racebattle.World::default(),
)

///|
let path : @racebattle.Path = @racebattle.Path::new([
  Vec2::new(68, 331),
  Vec2::new(93, 368),
  Vec2::new(107, 429),
  Vec2::new(207, 448),
  Vec2::new(240, 437),
  Vec2::new(284, 351),
  Vec2::new(309, 337),
  Vec2::new(384, 343),
  Vec2::new(442, 278),
  Vec2::new(429, 254),
  Vec2::new(374, 216),
  Vec2::new(310, 175),
])

///|
fn main {

}

///|
pub fn boot() -> Unit {
  let assets = Assets::load()
  racebattle_world.val = @racebattle.World::{
    field_img: assets.field,
    map_img: assets.racing_map,
    players: [
      @racebattle.Firefly::new(
        assets.firefly_sheet,
        assets.firefly_rev_sheet,
        Some(@firefly.get_me()),
        pos=Vec2::new(58, 315),
      ),
      //@racebattle.Player::new(assets.firefly_sheet, None),
      //@racebattle.Player::new(assets.firefly_sheet, None),
    ],
    camera: @racebattle.Camera::default(),
    me: @firefly.get_me(),
    path,
  }
}

///|
pub fn update() -> Unit {
  racebattle_world.val.update()
}

///|
pub fn render() -> Unit {
  racebattle_world.val.render()
}
