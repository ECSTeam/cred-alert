package db_test

import (
	"cred-alert/db"
	"encoding/json"
	"time"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"

	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RepositoryRepo", func() {
	var (
		repo     db.RepositoryRepository
		database *gorm.DB
		logger   lager.Logger
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("reporepo")
		var err error
		database, err = dbRunner.GormDB()
		Expect(err).NotTo(HaveOccurred())

		repo = db.NewRepositoryRepository(database)
	})

	Describe("FindOrCreate", func() {
		var (
			rawJSON      map[string]interface{}
			rawJSONBytes []byte
			repository   *db.Repository
		)

		BeforeEach(func() {
			rawJSON = map[string]interface{}{
				"path": "path-to-repo-on-disk",
				"name": "repo-name",
				"owner": map[string]interface{}{
					"login": "owner-name",
				},
				"private":        true,
				"default_branch": "master",
			}

			var err error
			rawJSONBytes, err = json.Marshal(rawJSON)
			Expect(err).NotTo(HaveOccurred())

			repository = &db.Repository{
				Name:          "repo-name",
				Owner:         "owner-name",
				Path:          "path-to-repo-on-disk",
				SSHURL:        "repo-ssh-url",
				Private:       true,
				DefaultBranch: "master",
				RawJSON:       rawJSONBytes,
			}
		})

		It("saves the repository to the database", func() {
			err := repo.FindOrCreate(repository)
			Expect(err).NotTo(HaveOccurred())

			savedRepository := &db.Repository{}
			database.Where("name = ? AND owner = ?", repository.Name, repository.Owner).Last(&savedRepository)

			Expect(savedRepository.Name).To(Equal("repo-name"))
			Expect(savedRepository.Owner).To(Equal("owner-name"))
			Expect(savedRepository.Path).To(Equal("path-to-repo-on-disk"))
			Expect(savedRepository.SSHURL).To(Equal("repo-ssh-url"))
			Expect(savedRepository.Private).To(BeTrue())
			Expect(savedRepository.DefaultBranch).To(Equal("master"))

			var actualRaw map[string]interface{}
			err = json.Unmarshal(savedRepository.RawJSON, &actualRaw)
			Expect(err).NotTo(HaveOccurred())

			Expect(actualRaw).To(Equal(rawJSON))
		})

		Context("when a repo with the same name and owner exists", func() {
			BeforeEach(func() {
				err := repo.FindOrCreate(&db.Repository{
					Name:          "repo-name",
					Owner:         "owner-name",
					Path:          "path-to-repo-on-disk",
					SSHURL:        "repo-ssh-url",
					Private:       true,
					DefaultBranch: "master",
					RawJSON:       rawJSONBytes,
				})
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns the saved repository", func() {
				err := repo.FindOrCreate(&db.Repository{
					Name:          "repo-name",
					Owner:         "owner-name",
					Path:          "path-to-repo-on-disk",
					SSHURL:        "repo-ssh-url",
					Private:       true,
					DefaultBranch: "master",
					RawJSON:       rawJSONBytes,
				})
				Expect(err).NotTo(HaveOccurred())

				var count int
				database.Model(&db.Repository{}).Where(
					"name = ? AND owner = ?", repository.Name, repository.Owner,
				).Count(&count)
				Expect(count).To(Equal(1))
			})
		})
	})

	Describe("FindOrCreate", func() {
		var (
			rawJSON      map[string]interface{}
			rawJSONBytes []byte
			repository   *db.Repository
		)

		BeforeEach(func() {
			rawJSON = map[string]interface{}{
				"path": "path-to-repo-on-disk",
				"name": "repo-name",
				"owner": map[string]interface{}{
					"login": "owner-name",
				},
				"private":        true,
				"default_branch": "master",
			}

			var err error
			rawJSONBytes, err = json.Marshal(rawJSON)
			Expect(err).NotTo(HaveOccurred())

			repository = &db.Repository{
				Name:          "repo-name",
				Owner:         "owner-name",
				Path:          "path-to-repo-on-disk",
				SSHURL:        "repo-ssh-url",
				Private:       true,
				DefaultBranch: "master",
				RawJSON:       rawJSONBytes,
			}
		})

		It("saves the repository to the database", func() {
			err := repo.Create(repository)
			Expect(err).NotTo(HaveOccurred())

			savedRepository := &db.Repository{}
			database.Where("name = ? AND owner = ?", repository.Name, repository.Owner).Last(&savedRepository)

			Expect(savedRepository.Name).To(Equal("repo-name"))
			Expect(savedRepository.Owner).To(Equal("owner-name"))
			Expect(savedRepository.Path).To(Equal("path-to-repo-on-disk"))
			Expect(savedRepository.SSHURL).To(Equal("repo-ssh-url"))
			Expect(savedRepository.Private).To(BeTrue())
			Expect(savedRepository.DefaultBranch).To(Equal("master"))

			var actualRaw map[string]interface{}
			err = json.Unmarshal(savedRepository.RawJSON, &actualRaw)
			Expect(err).NotTo(HaveOccurred())

			Expect(actualRaw).To(Equal(rawJSON))
		})
	})

	Describe("MarkAsCloned", func() {
		var repository *db.Repository

		BeforeEach(func() {
			repository = &db.Repository{
				Name:          "some-repo",
				Owner:         "some-owner",
				SSHURL:        "some-url",
				Private:       true,
				DefaultBranch: "some-branch",
				RawJSON:       []byte("some-json"),
			}
			err := repo.Create(repository)
			Expect(err).NotTo(HaveOccurred())
		})

		It("marks the repo as cloned", func() {
			err := repo.MarkAsCloned("some-owner", "some-repo", "some-path")
			Expect(err).NotTo(HaveOccurred())

			savedRepository := &db.Repository{}
			database.Where("name = ? AND owner = ?", repository.Name, repository.Owner).Last(&savedRepository)

			Expect(savedRepository.Cloned).To(BeTrue())
		})

		It("updates the path on the repo", func() {
			err := repo.MarkAsCloned("some-owner", "some-repo", "some-path")
			Expect(err).NotTo(HaveOccurred())

			savedRepository := &db.Repository{}
			database.Where("name = ? AND owner = ?", repository.Name, repository.Owner).Last(&savedRepository)

			Expect(savedRepository.Path).To(Equal("some-path"))
		})
	})

	Describe("NotFetchedSince", func() {
		var (
			savedFetch      db.Fetch
			savedRepository db.Repository
		)

		BeforeEach(func() {
			repository := &db.Repository{
				Name:          "some-repo",
				Owner:         "some-owner",
				SSHURL:        "some-url",
				Private:       true,
				DefaultBranch: "some-branch",
				RawJSON:       []byte("some-json"),
				Cloned:        true,
			}
			err := repo.Create(repository)
			Expect(err).NotTo(HaveOccurred())

			database.Where("name = ? AND owner = ?", repository.Name, repository.Owner).Last(&savedRepository)

			err = database.Model(&db.Fetch{}).Create(&db.Fetch{
				RepositoryID: savedRepository.ID,
				Changes:      []byte("changes"),
			}).Error
			Expect(err).NotTo(HaveOccurred())

			database.Where("repository_id = ?", repository.ID).Last(&savedFetch)
		})

		Context("when the repo's latest fetch is later than the given time", func() {
			BeforeEach(func() {
				t := time.Now().Add(-5 * time.Minute)
				database.Model(&db.Fetch{}).Where("id = ?", savedFetch.ID).Update("created_at", t)
			})

			It("does not return the repository", func() {
				repos, err := repo.NotFetchedSince(time.Now().Add(-10 * time.Minute))
				Expect(err).NotTo(HaveOccurred())
				Expect(repos).To(BeEmpty())
			})
		})

		Context("when the repo's latest fetch is not later than the given time", func() {
			BeforeEach(func() {
				t := time.Now().Add(-15 * time.Minute)
				database.Model(&db.Fetch{}).Where("id = ?", savedFetch.ID).Update("created_at", t)
			})

			It("returns the repo", func() {
				repos, err := repo.NotFetchedSince(time.Now().UTC().Add(-10 * time.Minute))
				Expect(err).NotTo(HaveOccurred())
				Expect(repos).To(ConsistOf(savedRepository))
			})

			Context("when the repository is disabled", func() {
				BeforeEach(func() {
					_, err := database.DB().Exec(`UPDATE repositories SET disabled = true WHERE id = ?`, savedRepository.ID)
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not return the repository", func() {
					repos, err := repo.NotFetchedSince(time.Now().Add(-10 * time.Minute))
					Expect(err).NotTo(HaveOccurred())
					Expect(repos).To(BeEmpty())
				})
			})

			Context("when the repository has not been cloned", func() {
				BeforeEach(func() {
					err := database.Model(&db.Repository{}).Where("id = ?", savedRepository.ID).Update("cloned", false).Error
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not return the repository", func() {
					repos, err := repo.NotFetchedSince(time.Now().Add(-10 * time.Minute))
					Expect(err).NotTo(HaveOccurred())
					Expect(repos).To(BeEmpty())
				})
			})
		})
	})

	Describe("NotScannedWithVersion", func() {
		var repository *db.Repository

		BeforeEach(func() {
			repository = &db.Repository{
				Name:          "some-repo",
				Owner:         "some-owner",
				SSHURL:        "some-url",
				Private:       true,
				DefaultBranch: "some-branch",
				RawJSON:       []byte("some-json"),
				Cloned:        true,
			}
			err := repo.Create(repository)
			Expect(err).NotTo(HaveOccurred())
		})

		It("returns an empty slice", func() {
			repos, err := repo.NotScannedWithVersion(42)
			Expect(err).NotTo(HaveOccurred())
			Expect(repos).To(BeEmpty())
		})

		Context("when the repository has scans for the specified version", func() {
			BeforeEach(func() {
				err := database.Create(&db.Scan{
					RepositoryID: &repository.ID,
				}).Error
				Expect(err).NotTo(HaveOccurred())
			})

			It("returns a slice with the repository", func() {
				repos, err := repo.NotScannedWithVersion(42)
				Expect(err).NotTo(HaveOccurred())
				Expect(repos).To(HaveLen(1))
				Expect(repos[0].Name).To(Equal("some-repo"))
				Expect(repos[0].Owner).To(Equal("some-owner"))
			})
		})
	})

	Describe("RegisterFailedFetch", func() {
		var repository *db.Repository

		BeforeEach(func() {
			repository = &db.Repository{
				Name:          "some-repo",
				Owner:         "some-owner",
				SSHURL:        "some-url",
				Private:       true,
				DefaultBranch: "some-branch",
				RawJSON:       []byte("some-json"),
				Cloned:        true,
				FailedFetches: db.FailedFetchThreshold - 2,
			}
		})

		JustBeforeEach(func() {
			err := repo.Create(repository)
			Expect(err).NotTo(HaveOccurred())
		})

		Context("when the repository has failed less than a threshold", func() {
			BeforeEach(func() {
				repository.FailedFetches = db.FailedFetchThreshold - 2
			})

			It("increments the failed fetch threshold", func() {
				err := repo.RegisterFailedFetch(logger, repository)
				Expect(err).NotTo(HaveOccurred())

				var failedFetches int
				err = database.DB().QueryRow(`
					SELECT failed_fetches
					FROM repositories
					WHERE id = ?
				`, repository.ID).Scan(&failedFetches)
				Expect(err).NotTo(HaveOccurred())

				Expect(failedFetches).To(Equal(db.FailedFetchThreshold - 1))
			})

			It("does not mark the repository as disabled", func() {
				err := repo.RegisterFailedFetch(logger, repository)
				Expect(err).NotTo(HaveOccurred())

				var disabled bool
				err = database.DB().QueryRow(`
					SELECT disabled
					FROM repositories
					WHERE id = ?
				`, repository.ID).Scan(&disabled)
				Expect(err).NotTo(HaveOccurred())

				Expect(disabled).To(BeFalse())
			})
		})

		Context("when the repository failing causes it to hit the threshold", func() {
			BeforeEach(func() {
				repository.FailedFetches = db.FailedFetchThreshold - 1
			})

			It("marks the repository as disabled", func() {
				err := repo.RegisterFailedFetch(logger, repository)
				Expect(err).NotTo(HaveOccurred())

				var disabled bool
				err = database.DB().QueryRow(`
					SELECT disabled
					FROM repositories
					WHERE id = ?
				`, repository.ID).Scan(&disabled)
				Expect(err).NotTo(HaveOccurred())

				Expect(disabled).To(BeTrue())
			})
		})

		It("returns an error when the repository can't be found", func() {
			err := repo.RegisterFailedFetch(logger, &db.Repository{
				Model: db.Model{
					ID: 1337,
				},
				Name:    "bad-repo",
				Owner:   "bad-owner",
				RawJSON: []byte("bad-json"),
			})
			Expect(err).To(HaveOccurred())
		})
	})
})
