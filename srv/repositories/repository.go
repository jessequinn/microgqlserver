package repositories

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	pb "github.com/jessequinn/microgqlserver/srv/proto/auth"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

// https://docs.mongodb.com/manual/reference/limits/#restrictions-on-db-names
// https://stackoverflow.com/questions/5916080/what-are-naming-conventions-for-mongodb
const (
	dbName           = "service"
	vesselCollection = "users"
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
			Id:       item.ID.String(),
			Name:     item.Name,
			Company:  item.Company,
			Email:    item.Email,
			Password: item.Password,
		})
	}
	//.All(&users); err != nil {
	//	return nil, err
	//}
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
	user := &pb.User{}
	if err := repo.collection().Find(bson.M{"email": email}).One(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *AuthRepository) Create(user *pb.User) error {
	if err := repo.collection().Insert(user); err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepository) Close() {
	repo.Session.Close()
}

// DB helper functions
func (repo *AuthRepository) collection() *mgo.Collection {
	return repo.Session.DB(dbName).C(vesselCollection)
}
