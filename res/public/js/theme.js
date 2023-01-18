$("document").ready(function () {
  let dark = "https://cdn.jsdelivr.net/npm/ashleycss@4.1.52/dist/themes/ashleycss-dark.min.css";
  let light = "https://cdn.jsdelivr.net/npm/ashleycss@4.1.52/dist/themes/ashleycss-sakura.min.css";

  let theme = Cookies.get("theme");

  if (!theme) Cookies.set("theme", "dark", { expires: 365 });

  if (theme == "dark") {
    $("#style").attr("href", dark);
    $("#theme").html("lights on!");
  } else {
    $("#style").attr("href", light);
    $("#theme").html("lights off!");
  }

  $("#theme").on("click", () => {
    let theme = Cookies.get("theme");
    if (theme == "dark") {
      Cookies.set("theme", "light", { expires: 365 });
      $("#style").attr("href", light);
      $("#theme").html("lights off!");
    } else {
      Cookies.set("theme", "dark", { expires: 365 });
      $("#style").attr("href", dark);
      $("#theme").html("lights on!");
    }
  });
});
