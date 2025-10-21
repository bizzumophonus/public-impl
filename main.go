package prettyapi

type PrettyType string

func (*PrettyType) awoo() string {
  return "I'm the public pretty type"
}

