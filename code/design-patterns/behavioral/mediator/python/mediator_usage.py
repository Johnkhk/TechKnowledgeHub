from mediator import StationManager, PassengerTrain, FreightTrain

station_manager = StationManager()

passenger_train = PassengerTrain(station_manager)
freight_train = FreightTrain(station_manager)

passenger_train.arrive()
freight_train.arrive()

passenger_train.depart()
freight_train.arrive()
freight_train.depart()
