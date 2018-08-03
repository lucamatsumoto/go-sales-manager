package gql

var Schema = `
	schema {
		query: Query
		# mutation: Mutation
	}
	type Query {
		hello(name: String!): String!
	}
	# type Mutation {}
` //replace this later
