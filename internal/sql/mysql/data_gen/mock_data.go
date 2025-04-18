package data_gen

import (
	"github.com/LucienVen/tech-backend/bootstrap"
	"github.com/LucienVen/tech-backend/internal/entity"
	"github.com/LucienVen/tech-backend/manager/log"
	"github.com/brianvoe/gofakeit/v7"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

const (
	CreatorMock string = "mock"
	UpdaterMock string = "mock"
)

func Mock() error {

	// 链接数据库
	//db := bootstrap.App().Mysql

	_ = gofakeit.Seed(0)

	// 创建老师
	//	学生
	//nowTime := time.Now().Unix()

	//insertSubjects()
	//InsertTeachers()
	// InsertStudent() // 需要关联班级
	return nil
}

// insertStudent 插入学生数据
func InsertStudent() {
	db := bootstrap.App().Mysql

	studentNames := GenerateBatchChineseNames(300)
	baseMode := entity.GenBaseModel(CreatorMock, UpdaterMock)

	class := []entity.Class{}
	err := db.Select(&class, "SELECT * FROM class WHERE is_delete = 0")
	if err != nil {
		log.Error("get class error", zap.Error(err))
	}

	students := make([]entity.Student, 0)

	classStudentNum := 10 // 每个班 10 人

	for times, item := range class {
		//fmt.Println(fmt.Sprintf("%d年%d班", item.GradeLevel, item.ClassNum))

		// 截取人名数组
		usedName := studentNames[times*classStudentNum : (times+1)*classStudentNum]
		//log.Info("usedName", zap.Any("usedName", usedName))

		for _, name := range usedName {

			gender := 1
			if !name.IsMale {
				gender = 2
			}

			students = append(students, entity.Student{
				Name:      name.Name,
				Gender:    int64(gender),
				ClassId:   item.Id,
				Phone:     gofakeit.Phone(),
				Email:     gofakeit.Email(),
				Passwd:    gofakeit.Password(true, true, true, true, false, 16),
				BaseModel: baseMode,
			})
		}

	}

	res, err := db.NamedExec(`INSERT INTO student (name, gender, class_id, phone, email, passwd, is_delete, creator, updater, create_time, update_time)
				VALUES (:name, :gender, :class_id, :phone, :email, :passwd, :is_delete, :creator, :updater, :create_time, :update_time)`, students)

	if err != nil {
		log.Error("insert student error", zap.Error(err))
	} else {
		log.Info("insert student success", zap.Any("students", students), zap.Any("res", res))

	}

}

// insertClasses 插入班级数据
func InsertClasses() {
	db := bootstrap.App().Mysql
	classNum := 5
	gradeLevel := 6

	baseMode := entity.GenBaseModel(CreatorMock, UpdaterMock)

	classData := make([]entity.Class, 0)

	for i := 1; i <= gradeLevel; i++ {
		for j := 1; j <= classNum; j++ {
			classData = append(classData, entity.Class{
				MainTeacherId: 0,
				ClassNum:      int64(j),
				GradeLevel:    int64(i),
				BaseModel:     baseMode,
			})
		}
	}

	res, err := db.NamedExec(`INSERT INTO class (class_num, grade_level, main_teacher_id, is_delete, creator, updater, create_time, update_time) 
							values (:class_num, :grade_level, :main_teacher_id, :is_delete, :creator, :updater, :create_time, :update_time)`, classData)

	if err != nil {
		log.Error("insert class error", zap.Error(err))
	} else {
		log.Info("insert class success", zap.Any("class", classData), zap.Any("res", res))
	}
}

// insertTeachers 插入老师数据
func InsertTeachers() {
	db := bootstrap.App().Mysql
	//teachers := []string{"张伟", "王伟", "王芳", "李伟", "王秀英", "李秀英", "李娜", "张秀英", "刘伟", "张敏"}
	teachers := GenerateBatchChineseNames(20)

	baseMode := entity.GenBaseModel(CreatorMock, UpdaterMock)

	teacherData := make([]entity.Teacher, len(teachers))
	for index, name := range teachers {
		gender := 1
		if !name.IsMale {
			gender = 2
		}
		teacherData[index] = entity.Teacher{
			Name:      name.Name,
			Age:       int64(gofakeit.Number(20, 60)),
			Gender:    int64(gender),
			Phone:     gofakeit.Phone(),
			Email:     gofakeit.Email(),
			Passwd:    gofakeit.Password(true, true, true, true, false, 16),
			BaseModel: baseMode,
		}
	}

	res, err := db.NamedExec(`INSERT INTO teacher (name, age, gender, phone, email, passwd, is_delete, creator, updater, create_time, update_time)
				VALUES (:name, :age, :gender, :phone, :email, :passwd, :is_delete, :creator, :updater, :create_time, :update_time)`, teacherData)

	if err != nil {
		log.Error("insert teacher error", zap.Error(err))
	} else {
		log.Info("insert teacher success", zap.Any("teachers", teacherData), zap.Any("res", res))
	}
}

// insertSubjects 插入学科数据
func insertSubjects() {
	db := bootstrap.App().Mysql
	subjects := []string{"语文", "数学", "英语", "科学", "体育"}
	timenow := time.Now().Unix()
	for _, name := range subjects {
		res, err := db.Exec("INSERT INTO subject (name, description, is_delete, creator, updater, create_time, update_time) VALUES (?, ?, 0, ?, ?, ?, ?)",
			name, gofakeit.Sentence(5), CreatorMock, UpdaterMock, timenow, timenow)
		if err != nil {
			log.Info("insert subject failed.", zap.String("name", name), zap.Error(err))
		} else {
			log.Info("insert subject.", zap.String("name", name), zap.Any("res", res))
		}

	}
}

// 考试
func InsertExam() {
	db := bootstrap.App().Mysql
	// 2024年，2023年，
	// 语数英，各 4 场，上下学期各 2

	// 获取学科 ID
	subjects := []entity.Subject{}
	err := db.Select(&subjects, "SELECT id, name FROM subject WHERE is_delete = 0 and name in ('语文', '数学', '英语')")
	if err != nil {
		log.Error("get subject err", zap.Error(err))
		return
	}

	//fmt.Println(subjects)

	baseMode := entity.GenBaseModel(CreatorMock, UpdaterMock)

	// 生成考试名称
	examName := make(map[string]int64, 0) // map[考试名称]学科 id
	for _, subject := range subjects {
		for i, s := range []string{"2023年上学期", "2023年下学期", "2024年上学期", "2024年下学期"} {
			examName[s+subject.Name+"考试"] = subject.Id

			if i <= 2 {
				// 2023年度

			} else {
				//2024年度

			}

		}
	}

	examData := make([]entity.Exam, 0)

	for en, sid := range examName {
		examData = append(examData, entity.Exam{
			Id:        0,
			Name:      en,
			SubjectId: sid,
			Year:      0,
			BaseModel: baseMode,
		})
	}

}
