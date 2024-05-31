package handler

import (
	"context"
	"database-example/dtos"
	"database-example/proto/encounter"
	"database-example/service"
)

type EncounterHandlergRPC struct {
	EncounterService service.IEncounterService
	encounter.UnimplementedEncounterServiceServer
}

func (handler *EncounterHandlergRPC) GetAll(ctx context.Context, in *encounter.Empty) (*encounter.EncounterDtoList, error) {
	allEncounters, err := handler.EncounterService.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert allEncounters to EncounterDtoList
	var encounterDtos []*encounter.EncounterDto

	for _, encounter1 := range allEncounters {
		encounterDto := ConvertDtosEncounterDtoToEncounterEncounterDto(&encounter1)

		encounterDtos = append(encounterDtos, encounterDto)
	}

	// Create and return EncounterDtoList
	return &encounter.EncounterDtoList{Encounters: encounterDtos}, nil

}

func (handler *EncounterHandlergRPC) Create(ctx context.Context, in *encounter.EncounterDto) (*encounter.EncounterDto, error) {

	encounterDto := ConvertEncounterEncounterDtoToDtosEncounterDto(in)

	createdEncounter, err := handler.EncounterService.Create(encounterDto)
	if err != nil {
		return nil, err
	}
	encounterResponse := ConvertDtosEncounterDtoToEncounterEncounterDto(createdEncounter)

	return encounterResponse, nil
}

func (handler *EncounterHandlergRPC) Update(ctx context.Context, in *encounter.EncounterDto) (*encounter.EncounterDto, error) {

	encounterDto := ConvertEncounterEncounterDtoToDtosEncounterDto(in)
	updatedEncounter, err := handler.EncounterService.Update(encounterDto)
	if err != nil {
		return nil, err
	}

	encounterResult := ConvertDtosEncounterDtoToEncounterEncounterDto(updatedEncounter)
	return encounterResult, nil
}

func (handler *EncounterHandlergRPC) Delete(ctx context.Context, in *encounter.DeleteEncounterRequest) (*encounter.Empty, error) {

	err := handler.EncounterService.Delete(in.GetId())
	if err != nil {
		return nil, err
	}

	empty := &encounter.Empty{}

	return empty, nil
}

func (handler *EncounterHandlergRPC) GetAllActive(ctx context.Context, in *encounter.Empty) (*encounter.EncounterDtoList, error) {
	encounters, err := handler.EncounterService.GetAllActive()
	if err != nil {
		return nil, err
	}

	var encounterDtos []*encounter.EncounterDto

	for _, encounter1 := range encounters {
		encounterDto := ConvertDtosEncounterDtoToEncounterEncounterDto(&encounter1)

		encounterDtos = append(encounterDtos, encounterDto)
	}

	return &encounter.EncounterDtoList{Encounters: encounterDtos}, nil
}

func ConvertDtosEncounterDtoToEncounterEncounterDto(dtosDto *dtos.EncounterDto) *encounter.EncounterDto {
	coordinatesResponse := &encounter.CoordinateDto{
		Longitude: dtosDto.Coordinates.Longitude,
		Latitude:  dtosDto.Coordinates.Latitude,
	}

	encounterResponse := &encounter.EncounterDto{
		Id:          dtosDto.ID,
		Name:        dtosDto.Name,
		Description: dtosDto.Description,
		Coordinates: coordinatesResponse,
		Xp:          dtosDto.Xp,
		Status:      encounter.EncounterStatus(dtosDto.Status),
		Type:        encounter.EncounterType(dtosDto.Type),
		Range:       dtosDto.Range,
	}

	if dtosDto.ImageUrl != nil {
		encounterResponse.ImageUrl = *dtosDto.ImageUrl
	}

	if dtosDto.MiscEncounterTask != nil {
		encounterResponse.MiscEncounterTask = *dtosDto.MiscEncounterTask
	}

	if dtosDto.SocialEncounterRequiredPeople != nil {
		encounterResponse.SocialEncounterRequiredPeople = *dtosDto.SocialEncounterRequiredPeople
	}

	if dtosDto.CheckpointId != nil {
		encounterResponse.CheckpointId = *dtosDto.CheckpointId
	}

	if dtosDto.IsRequired != nil {
		encounterResponse.IsRequired = *dtosDto.IsRequired
	}

	return encounterResponse
}

func ConvertEncounterEncounterDtoToDtosEncounterDto(dtoEncounterDto *encounter.EncounterDto) *dtos.EncounterDto {
	coordinatesDto := &dtos.CoordinateDto{
		Longitude: dtoEncounterDto.Coordinates.Longitude,
		Latitude:  dtoEncounterDto.Coordinates.Latitude,
	}

	encounterDto := &dtos.EncounterDto{
		ID:                            dtoEncounterDto.GetId(),
		Description:                   dtoEncounterDto.Description,
		Coordinates:                   *coordinatesDto,
		Xp:                            dtoEncounterDto.Xp,
		Status:                        dtos.EncounterStatus(dtoEncounterDto.Status),
		Type:                          dtos.EncounterType(dtoEncounterDto.Type),
		Range:                         dtoEncounterDto.Range,
		ImageUrl:                      &dtoEncounterDto.ImageUrl,
		MiscEncounterTask:             &dtoEncounterDto.MiscEncounterTask,
		SocialEncounterRequiredPeople: &dtoEncounterDto.SocialEncounterRequiredPeople,
		CheckpointId:                  &dtoEncounterDto.CheckpointId,
		IsRequired:                    &dtoEncounterDto.IsRequired,
	}

	return encounterDto
}
