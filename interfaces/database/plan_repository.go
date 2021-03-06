package database

import (
	"CoopeLunch-api/domain"
)

type PlanRepository struct {
	SqlHandler
}

func (repo *PlanRepository) All() (plans []domain.TPlan, err error) {
	rows, err := repo.Query(
		"SELECT * FROM plans",
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.TPlan

		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId, &p.ParticipantUsersCount,
		)
		if err != nil {
			panic(err.Error())
		}
		plans = append(plans, p)
	}

	return
}

func (repo *PlanRepository) GetByUserId(userId int) (plans []domain.TPlan, err error) {
	rows, err := repo.Query(
		"SELECT * FROM plans WHERE OwnerUserId = ?",
		userId,
	)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var p domain.TPlan

		err := rows.Scan(
			&p.ID, &p.ShopName, &p.MeetPlace, &p.MaxPeopleNumber, &p.MinPeopleNumber, &p.MeetTime, &p.PlanStatus, &p.OwnerUserId, &p.ParticipantUsersCount,
		)
		if err != nil {
			panic(err.Error())
		}
		plans = append(plans, p)
	}

	return
}

func (repo *PlanRepository) Insert(plan domain.TPlanInsert) (id int, err error) {
	exe, err := repo.Execute(
		"INSERT INTO plans(ShopName, MeetPlace, MaxPeopleNumber, MinPeopleNumber, MeetTime,  PlanStatus, OwnerUserId) VALUES(?, ?, ?, ?, ?, ?, ?)",
		plan.ShopName, plan.MeetPlace, plan.MaxPeopleNumber, plan.MinPeopleNumber, plan.MeetTime, plan.PlanStatus, plan.OwnerUserId,
	)
	if err != nil {
		return id, err
	}

	rawId, err := exe.LastInsertId()
	if err != nil {
		return id, err
	}
	id = int(rawId)

	_, err = repo.Execute(
		"INSERT INTO plan_participant_users (UserId, PlanId) VALUES(?, ?)",
		plan.OwnerUserId,
		id,
	)
	if err != nil {
		return id, err
	}

	return
}
