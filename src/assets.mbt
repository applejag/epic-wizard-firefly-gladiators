///|
struct Assets {
  field : @firefly.Image
  racing_map : @firefly.Image
  firefly_sheet : @util.SpriteSheet
  firefly_rev_sheet : @util.SpriteSheet
}

///|
fn Assets::load() -> Assets {
  Assets::{
    field: @firefly.load_file("field").unwrap().as_image(),
    racing_map: @firefly.load_file("racing-map").unwrap().as_image(),
    firefly_sheet: @util.SpriteSheet::split_image_by_count(
      @firefly.load_file("firefly").unwrap().as_image(),
      Size::new(7, 1),
    ),
    firefly_rev_sheet: @util.SpriteSheet::split_image_by_count(
      @firefly.load_file("firefly-rev").unwrap().as_image(),
      Size::new(7, 1),
    ),
  }
}
