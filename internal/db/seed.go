package db

import (
	"context"
	"fmt"
	"log"
	"math/rand"


	"github.com/everestp/Social_go/internal/store"
)

var usernames = []string{
	"pixelstorm", "novabright", "zephyrlane", "lunarspace", "kaizenwolf", "riovortex",
	"sageforest", "echomaster", "blazefirex", "onyxshadow", "vexmorning", "neonfusion",
	"jaxonrider", "remilights", "skylerwave", "cruzventure", "acehunter", "rivenstorm",
	"dashoracle", "lyricfinder", "zenithglow", "frostblaze", "stormchase", "ravenspark",
	"phoenixray", "wrenlegacy", "jetstreamx", "slatecloud", "quinnshadow", "ryzeimpact",
	"shadehunter", "tidebreaker", "vanceblade", "wispdreams", "xandermoon", "zetrixnova",
	"cypherlance", "orionquest", "luminaster", "faydenlight",
}

var postTitles = []string{
	"My day at the park", "Go programming tips", "Hello World!", "Random thoughts", "Traveling to Nepal",
	"Foodie adventures", "Learning React", "Daily motivation", "Fun with friends", "Tech news",
}

var postContents = []string{
	"Today I learned something new about Go.",
	"I tried a new restaurant and it was amazing!",
	"Just finished a project, feeling accomplished.",
	"Sharing some tips on programming.",
	"Random musings about life and work.",
}

var commentContents = []string{
	"Great post!", "I totally agree.", "Thanks for sharing.", "Interesting thoughts.", "Well written!",
	"Can you explain more?", "Loved this!", "Haha, so true.", "Nice!", "I have a question about this.",
}

// --------------------------------------------------
// Seed Database
// --------------------------------------------------
func Seed(store store.Storage) error {
	
	ctx := context.Background()

	users := generateUsers(30)

	// Insert users into DB
	for i, user := range users {
		if err := store.User.Create(ctx, user); err != nil {
			log.Println("Error creating user:", err)
			return err
		}
			log.Printf(" Creating user %s %s", i,user)
	}

	posts := generatePosts(40, users)
	for i, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("Error creating post:", err)
			return err
		}
			log.Printf(" Doing POsts %s %s", i,post)
	}

	comments := generateComments(50, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("Error creating comment:", err)
			return err
		}
		log.Printf(" Doing comment %s",comment)
	}
 log.Println("Seeding Complete")
	return nil
}

// --------------------------------------------------
// Generate Dummy Users
// --------------------------------------------------
func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		base := usernames[i%len(usernames)]

		users[i] = &store.User{
			Username: base + fmt.Sprintf("%d", i),
			Email:    fmt.Sprintf("%s%d@example.com", base, i),
			Password: "password123", // plain password for now
		}
	}

	return users
}

// --------------------------------------------------
// Generate Dummy Posts
// --------------------------------------------------
func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   postTitles[rand.Intn(len(postTitles))],
			Content: postContents[rand.Intn(len(postContents))],
			Tags:    []string{"fun", "random", "tech", "life"}[0:rand.Intn(3)+1],
		}
	}

	return posts
}

// --------------------------------------------------
// Generate Dummy Comments
// --------------------------------------------------
func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID:  posts[rand.Intn(len(posts))].ID,
			UserID:  users[rand.Intn(len(users))].ID,
			Content: commentContents[rand.Intn(len(commentContents))],
		}
	}

	return cms
}
