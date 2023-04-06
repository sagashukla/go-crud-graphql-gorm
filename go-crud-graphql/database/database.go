package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sagarshukla785/go-crud-graphql/graph/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"math/rand"
	"strconv"
)

var DBM *gorm.DB;

func ConnectToMyDB() *gorm.DB{
	dsn := "root:P9g78l@5b@/mydb"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DBM = db
	return DBM;
}

func GetJob(id string) *model.JobListing{
	jobId := id;
	var jobListing model.JobListing
	DBM.First(&jobListing, jobId)
	return &jobListing
}

func GetJobs() []*model.JobListing{
	var jobListings []*model.JobListing
	DBM.Find(&jobListings);
	return jobListings;
}

func CreateJobListing(jobInfo model.CreateJobListingInput) *model.JobListing{
	id := strconv.Itoa(rand.Intn(200))
	job := model.JobListing{
		ID: id,
		Title: jobInfo.Title,
		Description: jobInfo.Description,
		Company: jobInfo.Company,
		URL: jobInfo.URL,
	}
	inserted := DBM.Create(&job)
	if inserted.Error != nil {
		panic("Error while inserting the job");
	}

	return &job
}

func UpdateJobListing(jobId string, jobInfo model.UpdateJobListingInput) *model.JobListing{
	id := jobId
	job := model.JobListing{
		Title: jobInfo.Title,
		Description: jobInfo.Description,
		Company: jobInfo.Company,
		URL: jobInfo.URL,
	}
	
	DBM.First(&job, id)
	DBM.Model(&job).Updates(&job)
	return &job
}

func DeleteJobListing(jobId string) *model.DeleteJobResponse {
	id := jobId
	DBM.Delete(&model.DeleteJobResponse{}, id)
	return &model.DeleteJobResponse{DeleteJobID: jobId}
}