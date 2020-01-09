package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	pb "github.com/jessequinn/microgqlserver/srv/authsrv/proto/auth"
)

// https://docs.mongodb.com/manual/reference/limits/#restrictions-on-db-names
// https://stackoverflow.com/questions/5916080/what-are-naming-conventions-for-mongodb
const (
	dbName          = "service"
	usersCollection = "users"
)

type Repository interface {
	GetAll() ([]*pb.User, error)
	Get(id string) (*pb.User, error)
	Create(user *pb.User) error
	GetByEmail(email string) (*pb.User, error)
	Close()
}

type AuthRepository struct {
	Session *mgo.Session
}

// Mongo related type
type AuthUser struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	Company  string        `bson:"company"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
}

func (repo *AuthRepository) GetAll() ([]*pb.User, error) {
	item := AuthUser{}
	var users []*pb.User
	items := repo.collection().Find(bson.M{}).Iter()
	for items.Next(&item) {
		users = append(users, &pb.User{
			Id:       item.ID.Hex(),
			Name:     item.Name,
			Company:  item.Company,
			Email:    item.Email,
			Password: item.Password,
		})
	}
	if err := items.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (repo *AuthRepository) Get(id string) (*pb.User, error) {
	var user *pb.User
	if err := repo.collection().Find(bson.M{"_id": id}).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepository) GetByEmail(email string) (*pb.User, error) {
	item := AuthUser{}
	if err := repo.collection().Find(bson.M{"email": email}).One(&item); err != nil {
		return nil, err
	}
	return &pb.User{
		Id:       item.ID.Hex(),
		Name:     item.Name,
		Company:  item.Company,
		Email:    item.Email,
		Password: item.Password,
	}, nil
}

func (repo *AuthRepository) Create(user *pb.User) error {
	// Generate Object ID
	i := bson.NewObjectId()
	user.Id = i.Hex()
	if err := repo.collection().Insert(bson.M{
		"_id":      i,
		"name":     user.Name,
		"company":  user.Company,
		"email":    user.Email,
		"password": user.Password,
	}); err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepository) Close() {
	repo.Session.Close()
}

// DB helper functions
func (repo *AuthRepository) collection() *mgo.Collection {
	return repo.Session.DB(dbName).C(usersCollection)
}
