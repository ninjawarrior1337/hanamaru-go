{
  dockerTools,
  hanamaru,
  bash,
  cacert,
}:
dockerTools.buildLayeredImage {
  name = "hanamaru-go";
  tag = "latest";
  contents = [
    hanamaru
    bash
    cacert
  ];
  config = {
    Cmd = ["${hanamaru}/bin/hanamaru"];
    Env = [
      "IN_DOCKER=true"
    ];
    Volumes = {
      "/data" = {};
    };
  };
}
