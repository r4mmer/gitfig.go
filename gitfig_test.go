package gitfig

import "testing"

// type YamlConfig struct {
//   Server string `yaml:"Server"`
//   Foo string `yaml:"Foo"`
// }

type ShortJson struct {
  Glossary struct {
    Title string `json:"title"`;
    GlossDiv struct{
      Title string `json:"title"`;
      GlossList struct {
        GlossEntry struct {
          ID, SortAs, GlossTerm, Acronym, Abbrev, GlossSee string;
          GlossDef struct {
            Para string `json:"para"`;
            GlossSeeAlso []string;
          };
        };
      };
    };
  } `json:"glossary"`
}

// func TestYaml(t *testing.T) {
//   vy := YamlConfig{}
//   err := LoadYaml("prod/proj1/frontend.yaml", &vy, &FigOptions{
//     Repo: "https://github.com/r4mmer/config",
//     Branch: "proj",
//   })
//
//   if err != nil {
//     t.Errorf("%v", err.Error())
//   }
//   if vy.Foo != "bar" {
//     t.Errorf("did not get foo=bar...got=%v?", vy.Foo)
//   }
// }

func TestJson(t *testing.T) {
  v := ShortJson{}
  err := LoadJson("json/short.json", &v, &FigOptions{
    Repo: "https://github.com/git-fixtures/basic",
    Branch: "master",
  })
  if err != nil {
    t.Errorf("%v", err.Error())
  }
  if v.Glossary.Title != "example glossary" {
    t.Errorf("...?")
  }
}
