package prettyapi

type PrettyType string

func (*PrettyType) Awoo() string {
  return "I'm the public pretty type"
}

