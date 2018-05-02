package schema

const schema = `
schema {
	query: Query
	mutation: Mutation
}

type Query {
	me: Me
}

type PageInfo {
	startCursor: ID
	endCursor: ID
	hasNextPage: Boolean!
}

type Me {
	# current user ID
	id: ID!
	# email of logged in user
	email: String!
	# name of logged in user
	name: String
	# categories available to currently logged in user
	categories(first: Int,  after: String): CategoriesConnection!
	category(id: String!): Category
	card(id: String!): Card
}

type Category {
	id: ID!
	title: String
	cards(first: Int, after: String): CardsConnection
}

type CategoriesEdge {
	cursor: ID!
	node: Category
}

type CategoriesConnection {
	totalCount: Int!
	edges: [CategoriesEdge]
	pageInfo: PageInfo!
}

type Card {
	id: ID!
	title: String
}

type CardsEdge {
	cursor: ID!
	node: Card
}

type CardsConnection {
	totalCount: Int!
	edges: [CardsEdge]
	pageInfo: PageInfo!
}

type Mutation {
	addCard(input: CardInput!): Card
}

input CardInput {
	title: String!
}
`
